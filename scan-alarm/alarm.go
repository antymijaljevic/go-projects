package main

import (
	"fmt"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"image"
	_ "image/jpeg" // Add for JPEG decoding
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/gordonklaus/portaudio"
)

const (
	expectedBarcode  = "WAKEUP123"
	captureImagePath = "capture.jpg"
)

type AlarmApp struct {
	window          fyne.Window
	alarmTime       time.Time
	statusLabel     *widget.Label
	alarmActive     bool
	stopAlarmChan   chan bool
	soundPlaying    bool
	startStopButton *widget.Button
}

func main() {
	a := app.New()
	alarmApp := &AlarmApp{
		window:        a.NewWindow("Smart Alarm"),
		stopAlarmChan: make(chan bool),
	}

	alarmApp.createUI()
	alarmApp.window.Resize(fyne.NewSize(400, 300))
	alarmApp.window.ShowAndRun()
}

func (app *AlarmApp) createUI() {
	timeSelector := widget.NewButton("Set Alarm Time", app.showTimePicker)
	app.startStopButton = widget.NewButton("Start Alarm", app.toggleAlarm)
	app.startStopButton.Disable()
	app.statusLabel = widget.NewLabel("Status: Alarm not set")

	qrButton := widget.NewButton("Generate QR Code", app.generateQRCode)
	instructions := widget.NewLabel("1. Set alarm time\n2. Start alarm\n3. When alarm rings, scan QR code")

	app.window.SetContent(container.NewVBox(
		timeSelector,
		app.statusLabel,
		app.startStopButton,
		layout.NewSpacer(),
		qrButton,
		instructions,
	))
}

func (app *AlarmApp) showTimePicker() {
	now := time.Now()
	hourEntry := widget.NewEntry()
	minuteEntry := widget.NewEntry()

	timeDialog := dialog.NewCustom("Set Alarm Time", "OK",
		container.NewVBox(
			widget.NewLabel("Hour (00-23):"),
			hourEntry,
			widget.NewLabel("Minute (00-59):"),
			minuteEntry,
		),
		app.window,
	)

	timeDialog.SetOnClosed(func() {
		hour, _ := strconv.Atoi(hourEntry.Text)
		minute, _ := strconv.Atoi(minuteEntry.Text)
		app.alarmTime = time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, time.Local)
		if time.Now().After(app.alarmTime) {
			app.alarmTime = app.alarmTime.Add(24 * time.Hour)
		}
		app.startStopButton.Enable()
		app.statusLabel.SetText("Alarm set for: " + app.alarmTime.Format("2006-01-02 15:04"))
	})

	timeDialog.Show()
}

func (app *AlarmApp) toggleAlarm() {
	if app.alarmActive {
		app.stopAlarm()
	} else {
		app.startAlarm()
	}
}

func (app *AlarmApp) startAlarm() {
	app.alarmActive = true
	app.startStopButton.SetText("Stop Alarm")
	app.statusLabel.SetText("Alarm active - waiting for trigger...")

	go func() {
		time.Sleep(time.Until(app.alarmTime))
		if app.alarmActive {
			app.triggerAlarm()
		}
	}()
}

func (app *AlarmApp) stopAlarm() {
	app.alarmActive = false
	app.startStopButton.SetText("Start Alarm")
	app.statusLabel.SetText("Alarm stopped")
	app.stopAlarmChan <- true
}

func (app *AlarmApp) triggerAlarm() {
	app.window.RequestFocus()
	app.statusLabel.SetText("ALARM! Scan QR code to stop!")

	go func() {
		err := portaudio.Initialize()
		if err != nil {
			log.Printf("Audio error: %v", err)
			return
		}
		defer portaudio.Terminate()

		app.soundPlaying = true
		defer func() { app.soundPlaying = false }()

		for {
			select {
			case <-app.stopAlarmChan:
				return
			default:
				cmd := exec.Command("afplay", "alarm.wav")
				if err := cmd.Run(); err != nil {
					log.Printf("Error playing sound: %v", err)
				}
			}
		}
	}()

	scanButton := widget.NewButton("Scan QR Code Now", app.handleScan)
	app.window.SetContent(container.NewVBox(
		app.statusLabel,
		scanButton,
		widget.NewButton("Back", func() {
			app.createUI()
			app.stopAlarm()
		}),
	))
}

func (app *AlarmApp) handleScan() {
	err := captureCameraImage()
	if err != nil {
		dialog.ShowError(fmt.Errorf("Camera error: %v", err), app.window)
		return
	}

	value, err := decodeBarcodeFromImage()
	if err != nil {
		dialog.ShowError(fmt.Errorf("Scan failed: %v", err), app.window)
		return
	}

	if value == expectedBarcode {
		app.stopAlarm()
		app.createUI()
		dialog.ShowInformation("Success", "Alarm disabled!", app.window)
	} else {
		dialog.ShowError(fmt.Errorf("Invalid QR code"), app.window)
	}
}

func (app *AlarmApp) generateQRCode() {
	qrCode, _ := qr.Encode(expectedBarcode, qr.M, qr.Auto)
	qrImage, _ := barcode.Scale(qrCode, 300, 300)

	img := canvas.NewImageFromImage(qrImage)
	img.FillMode = canvas.ImageFillOriginal
	qrDialog := dialog.NewCustom("Your QR Code", "Close", img, app.window)
	qrDialog.Resize(fyne.NewSize(320, 320))
	qrDialog.Show()
}

func captureCameraImage() error {
	// Increase warmup time and add error handling
	cmd := exec.Command("imagesnap", "-w", "2", captureImagePath)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("camera failed: %v", err)
	}

	// Verify image was created
	if _, err := os.Stat(captureImagePath); os.IsNotExist(err) {
		return fmt.Errorf("no image captured")
	}

	return nil
}

func decodeBarcodeFromImage() (string, error) {
	file, err := os.Open(captureImagePath)
	if err != nil {
		return "", fmt.Errorf("failed to open image: %v", err)
	}
	defer file.Close()
	defer os.Remove(captureImagePath) // Clean up captured image after processing

	img, _, err := image.Decode(file)
	if err != nil {
		return "", fmt.Errorf("invalid image format: %v", err)
	}

	// Convert to grayscale for better QR detection
	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		return "", fmt.Errorf("image processing failed: %v", err)
	}

	// Try multiple barcode formats
	reader := qrcode.NewQRCodeReader()
	hints := map[gozxing.DecodeHintType]interface{}{
		gozxing.DecodeHintType_TRY_HARDER: true,
	}

	result, err := reader.Decode(bmp, hints)
	if err != nil {
		return "", fmt.Errorf("no QR code found: %v", err)
	}

	return result.GetText(), nil
}
