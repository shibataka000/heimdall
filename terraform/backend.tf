terraform {
  backend "s3" {
    bucket       = "sbtk-tfstate"
    key          = "heimdall"
    region       = "ap-northeast-1"
    use_lockfile = true
  }
}
