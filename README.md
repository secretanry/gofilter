# Filters App

The app is designed to apply 4 filters:
1) Grayscale
2) Mirroring
3) Resize
4) Blur

App works according to pipes-and-filters pattern using Golang.
Every filter applied in separate goroutine and transfers the data to next filter process using go channels.

## How to run
1) You need to have opencv4 installed on your machine and exported to pkg-config
2) Install gocv.io/x/gocv package
3) Run the app using go run .