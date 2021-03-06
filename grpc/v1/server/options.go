package gigrpc

import giconfig "github.com/b2wdigital/goignite/v2/config"

type Options struct {
	Port                 int
	MaxConcurrentStreams int64
	TLS                  struct {
		Enabled  bool
		CertFile string
		KeyFile  string
		CAFile   string `config:"CAFile"`
	} `config:"tls"`
}

func DefaultOptions() (*Options, error) {

	o := &Options{}

	err := giconfig.UnmarshalWithPath(root, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}
