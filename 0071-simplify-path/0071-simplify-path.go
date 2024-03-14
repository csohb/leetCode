func simplifyPath(path string) string {
    paths := make([]string, 0)
    dirs := strings.Split(path, "/")

    for _, dir := range dirs {
        //fmt.Println("dir:",dir)
        if dir == "" {
            continue
        }

        if dir == "." {
            continue
        }

        if dir != ".." {
            paths = append(paths, dir)
        }

        if dir == ".." && len(paths) > 0 {
            paths = paths[:len(paths)-1]
        }
    }

    return "/" + strings.Join(paths, "/")
}