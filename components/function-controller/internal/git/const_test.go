package git_test

const (
	testSSHPrivateKey = `
-----BEGIN OPENSSH PRIVATE KEY-----
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
-----END OPENSSH PRIVATE KEY-----`

	testSSHPrivateKeyPassphrase = `-----BEGIN OPENSSH PRIVATE KEY-----
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
bJOkA627VTiLKjoJTfnnYQ6KcnM=
-----END OPENSSH PRIVATE KEY-----`
)
