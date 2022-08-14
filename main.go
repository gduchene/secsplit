// SPDX-FileCopyrightText: © 2021 Grégoire Duchêne <gduchene@awhk.org>
// SPDX-License-Identifier: ISC

package main

import (
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/hashicorp/vault/shamir"
)

var (
	combine = flag.Bool("c", false, "combine instead of split")
	parts   = flag.Int("p", 10, "parts to split the secret into")
	thres   = flag.Int("t", 2, "threshold required to reconstruct the secret")
)

func main() {
	flag.Parse()

	buf, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to read the secret:", err)
		os.Exit(1)
	}

	if !*combine {
		parts, err := shamir.Split(buf, *parts, *thres)
		if err != nil {
			fmt.Fprintln(os.Stderr, "failed to split the secret:", err)
			os.Exit(1)
		}
		for _, part := range parts {
			fmt.Println(base64.StdEncoding.EncodeToString(part))
		}
		return
	}

	parts := [][]byte{}
	for _, buf := range bytes.Split(buf, []byte("\n")) {
		if len(buf) == 0 {
			continue
		}
		part := make([]byte, base64.StdEncoding.DecodedLen(len(buf)))
		n, err := base64.StdEncoding.Decode(part, buf)
		if err != nil {
			fmt.Fprintln(os.Stderr, "failed to decode a part:", err)
			os.Exit(1)
		}
		parts = append(parts, part[:n])
	}
	sec, err := shamir.Combine(parts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to combine parts:", err)
		os.Exit(1)
	}
	fmt.Print(string(sec))
}
