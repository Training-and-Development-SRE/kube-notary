/*
 * Copyright (c) 2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package verify

import (
	"github.com/vchain-us/kube-notary/pkg/image"
	"github.com/vchain-us/vcn/pkg/api"
)

// ImageID returns the hast string and the BlockchainVerification for the given imageID
func ImageID(imageID string, options ...Option) (hash string, verification *api.BlockchainVerification, err error) {
	o, err := makeOptions(options...)
	if err != nil {
		return
	}

	digest, err := image.Resolve(imageID, o.keychain)
	if err != nil {
		return
	}
	return getVerification(digest, o)
}
