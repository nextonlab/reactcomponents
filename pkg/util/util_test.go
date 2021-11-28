package util

//func TestWget(t *testing.T) {
//	var (
//		name string
//		err  error
//	)
//	for i := 0; i < 2; i++ {
//		name, err = Wget(context.Background(), "n1", "https://raw.githubusercontent.com/pingcap/tikv/master/Cargo.toml", ".")
//		if err != nil {
//			t.Fatalf("download failed %v", err)
//		}
//
//		if !IsFileExist(context.Background(), "n1", name) {
//			t.Fatalf("stat %s failed %v", name, err)
//		}
//	}
//
//	RemoveDir(context.Background(), "n1", name)
//}
//
//func TestInstallArchive(t *testing.T) {
//	tmpDir := fmt.Sprintf("/tmp/chaos/test_%d", time.Now().UnixNano())
//	t.Logf("install on %s", tmpDir)
//
//	Mkdir(context.Background(), "n1", tmpDir)
//	defer RemoveDir(context.Background(), "n1", tmpDir)
//
//	err := InstallArchive(context.Background(), "n1", "https://github.com/pingcap/tipocket/archive/master.zip", path.Join(tmpDir, "1"))
//	if err != nil {
//		t.Fatalf("install archive failed %v", err)
//	}
//
//	err = InstallArchive(context.Background(), "n1", "https://github.com/pingcap/tipocket/archive/master.tar.gz", path.Join(tmpDir, "2"))
//	if err != nil {
//		t.Fatalf("install archive failed %v", err)
//	}
//
//	archFile := path.Join(tmpDir, "a.tar.gz")
//	testCreateArchive(context.Background(), t, path.Join(tmpDir, "test"), archFile)
//	err = InstallArchive(context.Background(), "n1", "file://"+archFile, path.Join(tmpDir, "3"))
//	if err != nil {
//		t.Fatalf("install archive failed %v", err)
//	}
//}
//
//func testCreateArchive(ctx context.Context, t *testing.T, srcDir string, name string) {
//	t.Logf("crate archieve %s from %s", name, srcDir)
//	Mkdir(ctx, "n1", srcDir)
//	WriteFile(ctx, "n1", path.Join(srcDir, "a.log"), "\"hello world\"")
//
//	if err := ssh.Exec(ctx, "n1", "tar", "-cf", name, "-C", srcDir, "."); err != nil {
//		t.Fatalf("tar %s to %s failed %v", srcDir, name, err)
//	}
//}
//
//func TestDaemon(t *testing.T) {
//	t.Log("test may only be run in the chaos docker")
//
//	tmpDir := fmt.Sprintf("/tmp/chaos/var_%d", time.Now().UnixNano())
//	Mkdir(context.Background(), "n1", tmpDir)
//	defer RemoveDir(context.Background(), "n1", tmpDir)
/