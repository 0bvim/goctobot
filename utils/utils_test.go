package utils

// func TestColorized(t *testing.T) {
// 	tests := []struct {
// 		color  string
// 		text   string
// 		expect string
// 	}{
// 		{Red, "Hello", Red + "Hello" + Reset},
// 		{Green, "World", Green + "World" + Reset},
// 	}
//
// 	for _, tt := range tests {
// 		t.Run(fmt.Sprintf("Colorize(%q, %q)", tt.color, tt.text), func(t *testing.T) {
// 			result := Colorize(tt.color, tt.text)
// 			if result != tt.expect {
// 				t.Errorf("Colorize(%q, %q) = %q; want %q", tt.color, tt.text, result, tt.expect)
// 			}
// 		})
// 	}
// }

// func TestGetToken(t *testing.T) {
// 	// Save the original value of the environment variable to restore later
// 	originalToken := os.Getenv("personal_github_token")
// 	defer os.Setenv("personal_github_token", originalToken)
//
// 	// Test with an empty token
// 	os.Setenv("personal_github_token", "")
//
// 	// Capture the output of GetToken
// 	output := captureOutput(func() {
// 		defer func() {
// 			if r := recover(); r != nil {
// 				// Expected to exit, so we ignore the panic
// 			}
// 		}()
// 		GetToken()
// 	})
//
// 	expectedOutput := Colorize(Red, "Error: 'personal_github_token' environment variable not set.") + "\n" +
// 		Colorize(Magenta, "To resolve this: ") + "\n" +
// 		Colorize(Green, "      1. Generate a GitHub personal access token with the 'user:follow' and 'read:user' scopes at https://github.com/settings/tokens.\n"+
// 			"      2. Set the token in your environment with:\n"+
// 			"      export personal_github_token=\"your_token_here\"\n"+
// 			"      3. To make this change permanent, add it to your '~/.bashrc' with:\n"+
// 			"      echo 'export personal_github_token=\"your_token_here\"' >> ~/.bashrc\n"+
// 			"      source ~/.bashrc\n\n"+
// 			"      After setting up the token, you can run OctoBot commands with:\n"+
// 			"      ghbot <command> [username]\n\n"+
// 			"      For more details, visit the GitHub repository.") + "\n"
//
// 	if output != expectedOutput {
// 		t.Errorf("GetToken() output = %q; want %q", output, expectedOutput)
// 	}
// }

// Helper function to capture output
// func captureOutput(f func()) string {
// 	var buf bytes.Buffer
// 	writer := &buf
//
// 	// Redirect stdout to capture output
// 	old := os.Stdout
// 	defer func() { os.Stdout = old }()
// 	os.Stdout = writer.Bytes()
//
// 	f()
// 	return buf.String()
// }

// func captureOutput(f func()) string {
// 	var buf bytes.Buffer
// 	old := os.Stdout
// 	_, w, _ := os.Pipe()
// 	os.Stdout = w
//
// 	defer func() {
// 		w.Close()
// 		os.Stdout = old
// 	}()
//
// 	f()
//
// 	outC := make(chan []byte, 1)
// 	go func() {
// 		outC <- buf.Bytes()
// 	}()
//
// 	return string(<-outC)
// }
