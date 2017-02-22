package tda

type Lista struct {
	Primero Nodo
	Ultimo  Nodo
	Tam     int
}

func (this *Lista) Vacio() bool {
	if this.Tam == 0 {
		return true
	} else {
		return false
	}
}

func (this *Lista) Agregar_Primero(dato string) {
	if this.Vacio() {
		this.Primero.Dato = dato
		this.Ultimo.Dato = dato
		this.Primero = this.Ultimo
	} else {
		var tem Nodo
		tem = this.Primero
		// this.primero.dato = fact
		tem.Siguiente = &this.Primero
	}

}
