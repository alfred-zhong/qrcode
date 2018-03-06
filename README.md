# qrcode [![Build Status](https://www.travis-ci.org/alfred-zhong/qrcode.svg?branch=master)](https://www.travis-ci.org/alfred-zhong/qrcode) [![Go Report Card](https://goreportcard.com/badge/github.com/alfred-zhong/qrcode)](https://goreportcard.com/report/github.com/alfred-zhong/qrcode)

Tool to create qrcode image(.png) with text

Usage: qrcode [-f filepath] [-n] text

    -f  Path of file to create the image. File will be created in a temporary
        directory if not specified.
    
    -n  Don not open file after file created. Default file will be opened. 
        (Now only support macOSX)
