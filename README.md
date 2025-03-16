# Photo Viewer

Photo Viewer is a simple web application for displaying photos. This application is written in Go and supports gallery display and individual photo viewing.

## Purpose

This tool was developed to transfer photo data from a PC to a smartphone. Access the web application from your smartphone and download the photos to your smartphone.

## Features

- Gallery display of photos
- Individual photo viewing
- Photo download

## Setup

### Prerequisites

- Docker installed

### Build and Run

1. Clone the repository.

    ```sh
    git clone <repository-url>
    cd <repository-directory>
    ```

2. Build the Docker image.

    ```sh
    docker build -t photo-viewer .
    ```

3. Start the container.

    ```sh
    docker run -p 8080:8080 -v $(pwd)/pictures:/app/pictures photo-viewer
    ```

4. Access `http://localhost:8080` in your browser.

## Usage

- Access the main page to see the photos in the `pictures` directory displayed as a gallery.
- Click on a photo to go to the individual photo view page.
- On the individual photo view page, you can download the photo.