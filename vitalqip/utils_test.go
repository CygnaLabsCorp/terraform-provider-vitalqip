package vitalqip

func testAccConfigWithProvider(config string) string {
	return server + "\n" + config
}
