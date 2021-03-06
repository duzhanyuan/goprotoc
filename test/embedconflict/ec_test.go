// Copyright (c) 2013, Vastech SA (PTY) LTD. All rights reserved.
// http://code.google.com/p/gogoprotobuf/gogoproto
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package embedconflict

import (
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/dropbox/goprotoc/test_config"
)

func TestEmbedConflict(t *testing.T) {
	cmd := exec.Command("protoc", "--dgo_out=.", "-I="+config.ProtoPath, "ec.proto")
	data, err := cmd.CombinedOutput()
	if err == nil {
		t.Errorf("Expected error")
		if err := os.Remove("ec.pb.go"); err != nil {
			panic(err)
		}
	}
	t.Logf("received expected error = %v and output = %v", err, string(data))
}

func TestEmbedMarshaler(t *testing.T) {
	cmd := exec.Command("protoc", "--dgo_out=.", "-I="+config.ProtoPath, "em.proto")
	data, err := cmd.CombinedOutput()
	dataStr := string(data)
	t.Logf("received error = %v and output = %v", err, dataStr)
	if !strings.Contains(dataStr, "WARNING: found non-") || !strings.Contains(dataStr, "marshaler") {
		t.Errorf("Expected WARNING: found non-[marshaler] C with embedded marshaler D")
	}
	if err := os.Remove("em.pb.go"); err != nil {
		panic(err)
	}
}

func TestEmbedExtend(t *testing.T) {
	cmd := exec.Command("protoc", "--dgo_out=.", "-I="+config.ProtoPath, "ee.proto")
	data, err := cmd.CombinedOutput()
	if err == nil {
		t.Errorf("Expected error")
		if err := os.Remove("ee.pb.go"); err != nil {
			panic(err)
		}
	}
	t.Logf("received expected error = %v and output = %v", err, string(data))
}

func TestCustomName(t *testing.T) {
	cmd := exec.Command("protoc", "--dgo_out=.", "-I="+config.ProtoPath, "en.proto")
	data, err := cmd.CombinedOutput()
	if err == nil {
		t.Errorf("Expected error")
		if err := os.Remove("en.pb.go"); err != nil {
			panic(err)
		}
	}
	t.Logf("received expected error = %v and output = %v", err, string(data))
}

func TestRepeatedEmbed(t *testing.T) {
	cmd := exec.Command("protoc", "--dgo_out=.", "-I="+config.ProtoPath, "er.proto")
	data, err := cmd.CombinedOutput()
	if err == nil {
		t.Errorf("Expected error")
		if err := os.Remove("er.pb.go"); err != nil {
			panic(err)
		}
	}
	dataStr := string(data)
	t.Logf("received error = %v and output = %v", err, dataStr)
	warning := "ERROR: found repeated embedded field B in message A"
	if !strings.Contains(dataStr, warning) {
		t.Errorf("Expected " + warning)
	}
}
