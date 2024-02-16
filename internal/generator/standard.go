/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package generator

import (
	"github.com/lucasloureiror/AegisPass/internal/cli"
	"github.com/lucasloureiror/AegisPass/internal/shuffle"
	"sync"
)

type standard struct{}

func (standard) generate(input *cli.Input) (string, int, error) {

	var wg sync.WaitGroup
	var credits int
	wg.Add(4)
	upper := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	special := []byte("!@#$%&*-_")
	nums := []byte("0123456789")

	go func() {
		defer wg.Done()
		shuffle.Byte(&upper)
	}()
	go func() {
		defer wg.Done()
		shuffle.Byte(&special)
	}()

	go func() {
		defer wg.Done()
		shuffle.Byte(&nums)
	}()

	go func() {
		defer wg.Done()
		shuffle.Byte(&input.CharSet)
	}()

	wg.Wait()

	requirements := string(upper[0]) + string(special[0]) + string(nums[0])

	generated := shuffle.BuildString(input.CharSet, input.Size-3)

	generated = generated + requirements
	shuffle.String(&generated)

	return generated, credits, nil
}
