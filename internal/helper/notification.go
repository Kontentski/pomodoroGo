package helper

import (
	_"embed"
	"fmt"
	"os"
	"os/exec"
)

//go:embed notification.mp3
var notification []byte

// plays the notification sound when timer is up.
func PlayNotification(muted bool) error {
	if muted {
		return nil // No need to play anything when muted.
        }
	
	tempFileName := os.TempDir() + "/notification.mp3"

	// Check if the file already exists
	if _, err := os.Stat(tempFileName); os.IsNotExist(err) {
		// Create the notification sound file
		tempFile, err := os.Create(tempFileName)
		if err != nil {
			return fmt.Errorf("failed to create temp file: %w", err)
		}
		defer os.Remove(tempFile.Name()) // Clean up the temp file afterwards

		// Write the notification sound to the temp file
		if _, err := tempFile.Write(notification); err != nil {
			return fmt.Errorf("failed to write notification sound: %w", err)
		}

	}

	// Play the notification sound using an external player
	cmd := exec.Command("ffplay", "-nodisp", "-autoexit", tempFileName)
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start playback: %w", err)
	}
	// Wait for the command to finish
	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("error waiting for playback to finish: %w", err)
	}

	return nil
}

