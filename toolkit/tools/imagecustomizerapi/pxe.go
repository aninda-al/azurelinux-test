// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package imagecustomizerapi

import (
	"fmt"
	"strings"
)

var PxeIsoDownloadProtocols = []string{"http://"}

// Iso defines how the generated iso media should be configured.
type Pxe struct {
	IsoImageUrl string `yaml:"isoImageUrl"`
}

func (i *Pxe) IsValid() error {
	if i.IsoImageUrl != "" {
		protocolFound := false
		for _, protocol := range PxeIsoDownloadProtocols {
			if strings.HasPrefix(i.IsoImageUrl, protocol) {
				protocolFound = true
				break
			}
		}
		if !protocolFound {
			return fmt.Errorf("invalid iso image URL prefix in (%s). One of (%v) is expected.", i.IsoImageUrl, PxeIsoDownloadProtocols)
		}
	}
	return nil
}