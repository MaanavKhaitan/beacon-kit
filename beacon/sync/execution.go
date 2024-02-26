// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package sync

import (
	"context"
)

// CheckELSync checks if the execution layer is syncing.
func (s *Service) CheckELSync(ctx context.Context) error {
	// Call the ethClient to get the sync progress
	progress, err := s.ethClient.SyncProgress(ctx)
	if err != nil {
		return err
	}

	// Exit early if the node does not return a progress.
	// This means the node is in sync at the eth1 layer.
	if progress == nil {
		return nil
	}

	s.Logger().Warn(
		"execution client is attemping to sync.... ",
		"current_eth1", progress.CurrentBlock,
		"highest_eth1", progress.HighestBlock,
		"starting_eth1", progress.StartingBlock,
	)

	return ErrExecutionClientIsSyncing
}