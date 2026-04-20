const diagram_url = new URL("https://viewer.diagrams.net/?tags=%7B%7D&lightbox=1&highlight=0000ff&edit=_blank&layers=1&nav=1&title=Run-Length%20Encoding&dark=auto#Uhttps%3A%2F%2Fdrive.google.com%2Fuc%3Fid%3D1j8jntY4htsET-ArOSU-JDVIBnANmdVTv%26export%3Ddownload")

const open_diagram_button = document.querySelector('input#open_diagram_button');

open_diagram_button.addEventListener('click', e => {
	console.log("Opening diagram...", e);
	window.open(diagram_url, '_blank');
});

