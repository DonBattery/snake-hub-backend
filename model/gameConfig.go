package model

// GameConfig contains the rules of the game world
type GameConfig struct {
	Maps         []string `yaml:"maps"           mapstructure:"maps"`
	MaxConn      int      `yaml:"max_conn"       mapstructure:"max_conn"`
	ConnPerWorld int      `yaml:"conn_per_world" mapstructure:"conn_per_world"`
	RoundScore   int      `yaml:"round_score"    mapstructure:"round_score"`
}
