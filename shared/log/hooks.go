package log

import "github.com/sirupsen/logrus"

type EnvFieldHook struct {
	service string
	env     string
}

func NewEnvFieldHook(service string, env string) *EnvFieldHook {
	return &EnvFieldHook{
		service: service,
		env:     env,
	}
}

func (h *EnvFieldHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *EnvFieldHook) Fire(entry *logrus.Entry) error {
	entry.Data["service"] = h.service
	entry.Data["env"] = h.env

	return nil
}
