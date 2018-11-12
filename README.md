# deb822
```go
import (
    "fmt"

    "gitlab.com/kimtinh/godebian/deb822"
)

func main() {
    // Find paragraph which has Package field is vim
    vim := deb822.FindOne("a_Sources_file","Package","vim")

    // Find all paragraphs which have Format field is 1.0
    format1 := deb822.FindAll("a_Source_file","Format","1.0")

    // Find the first 10 paragraphs which
    anything := deb822.Find("a_Source_file","","",10)

    // Show all files of source package nano
    nano := deb822.FindOnePackage("a_Sources_file","Package","nano")
    for _, file := range nano.Files {
        fmt.Println(file["name"])
        fmt.Println(file["md5sum"])
    }
}
```

# control
```go
import "gitlab.com/kimtinh/godebian/control"

func main() {
    // Read debian/control
    ctrl := control.Parse("debian/control")
}
```

#  
<a href='https://ko-fi.com/W7W5KWLN' target='_blank'><img height='36' style='border:0px;height:36px;' src='https://az743702.vo.msecnd.net/cdn/kofi2.png?v=0' border='0' alt='Buy Me a Coffee at ko-fi.com' /></a>