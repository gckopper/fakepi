$name = "tester"
$arches = "amd64", "386", "arm64", "arm"
$oses = "darwin", "windows", "linux", "android"
$original_arch = $env:GOARCH
$original_os = $env:GOOS

foreach ($arch in $arches) {
    $env:GOARCH = $arch
    foreach ($os in $oses) {
        $env:GOOS = $os
        if ($os -eq "windows") {
            $os = "windows.exe"
        }
        go build -o "build/$name-$arch-$os" main.go
    }
}

$env:GOOS = $original_os
$env:GOARCH = $original_arch