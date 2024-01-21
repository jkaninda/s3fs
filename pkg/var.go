package pkg

const s3MountPath string = "/s3mnt"
const s3fsPasswdFile string = "/etc/passwd-s3fs"

var (
	accessKey  = ""
	secretKey  = ""
	bucketName = ""
	s3Endpoint = ""
	keep       = false
)
