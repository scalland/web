package updater

// S3Updater upgrades via S3/B2 bucket.
type S3Updater struct {
	Bucket    string
	Region    string
	Prefix    string
	AccessKey string
	SecretKey string
}
