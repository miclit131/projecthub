in temp.go td := WebGl{80, 8085, "index.html", "fields"} mit den 
entsprechenden Feldern befüllen und laufen lassen um die docker,
docker-compose und webgl.conf Datei zu erhalten und dann eine
Ordnerstruktur wie im Docker Ordner erstellen um ein WebGl Docker
image zu bekommen.

Das Templating wurde nur am Anfang gecoded wurde aber verworfen aufgrund
von anderen Prioritäten innerhalb des Projekts. Hier nochmal hochgeladen
für die nächste Gruppe die daran arbeitet.

type WebGl struct {
	InternalPort int
	ExternalPort int
	Index        string
	ConName      string
}

