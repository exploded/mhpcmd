// Sent command to USB via HID. Uses library from github.com/karalabe/usb
//

package main

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/karalabe/usb"
)

const mhpVendorID = 0x12bf
const mhpProductID = 0xff03
const mhpMessageSize = 8

func hidSend(message int) (err error) {
	// Enumerate all the HID devices matching the MHP ID
	hids, err := usb.EnumerateHid(mhpVendorID, mhpProductID)
	if err != nil {
		err = errors.New("mount hub pro not found")
		return
	}

	if len(hids) < 1 {
		err = errors.New("mount hub pro not found")
		return
	}

	if len(hids) > 1 {
		err = errors.New("only 1 Mount hub pro can be connected at the same time")
		return
	}

	myhid := hids[0]
	bs := make([]byte, mhpMessageSize)
	binary.LittleEndian.PutUint32(bs, uint32(message))
	mydevice, err := myhid.Open()
	if err != nil {
		return
	}
	defer mydevice.Close()
	i, err := mydevice.Write(bs)
	if err != nil {
		return
	}
	fmt.Println("Success.", i, "bytes written.")
	return
}
