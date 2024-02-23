#!/bin/bash

# Relative to parent dir
cd "`dirname "$0"`"/..

# Tailwind CSS
./resources/tailwindcss build -i css/styles.css -o public/tailwind.css --minify
