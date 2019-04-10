package classpath

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func Parse(jreOption, cpOption string,) *Classpath {

}

func (self *Classpath) ReadClass(className string) ([] byte, Entry, error) {

}

func (self *Classpath) String() string {

}
