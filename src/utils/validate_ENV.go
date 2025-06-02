package utils

func ValidationENV(env string, envDefault string) string {
    if env == "" {
        return envDefault
    }
    return env
}