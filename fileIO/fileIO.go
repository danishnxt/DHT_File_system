package fileIO

import (
	"DHT_NXT/consts"
	"DHT_NXT/util"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// have your file handling functions here

// Read file into multiple .dat slices into slice_directory
func CreateFileSlices(filename string, slice_directory string) {
	f_read, err := os.Open(filename)
	util.CheckError("", err)
	defer f_read.Close()
	slice_count := 0
	b1 := make([]byte, consts.FILE_BUF) // set this to > filesize to load in one go (lol)
	// clean folder and regen
	os.Remove(slice_directory)
	os.Mkdir(slice_directory, 0700)
	for {
		n1, err := f_read.Read(b1)
		write_buf := b1[:n1] // slice to only the data read
		if n1 == 0 {
			break
		}
		util.CheckError("", err)
		fmt.Printf("%d** Reading bytes from %s \n", n1, filename)
		f, err := os.Create(slice_directory + "/" + filename + ".nxt_dat" + strconv.Itoa(slice_count))
		util.CheckError("", err)
		f.Write(write_buf)
		f.Close()
		slice_count++
	}
}

// Recreate file from multiple .dat slices in slice_directory to single file
func ReadSliceRecreate(filename string, slice_directory string) {
	remade_file, err := os.Create(filename)
	util.CheckError("", err)
	entries, err := ioutil.ReadDir(slice_directory)
	util.CheckError("", err)
	fmt.Println("** FILE SAVE => Read file count: ", len(entries))
	for i := 0; i < len(entries); i++ {
		slice_name := entries[i].Name()
		slice_size := entries[i].Size()
		// another possible - read all into mem and homogenise - perf diff?
		buf := make([]byte, slice_size)
		read_slice, err := os.Open(slice_directory + slice_name)
		util.CheckError("", err)
		read_slice.Read(buf)
		_, err = remade_file.Write(buf)
		util.CheckError("", err)
	}
	remade_file.Close()
}
