package models

const rowTemplate = `
<div class="row">
	%s
</div>
`
const treasureTemplate = `
	<div class="column" style="background-color:#aaa; border-style: solid; border-color: coral;">
    	<p>Treasure</p>
  	</div>
`

const emptySlotTemplate = `
	<div class="column" style="border-style: solid; border-color: coral;">
    	<p></p>
  	</div>
`
const PlayerTemplate = `
	<div class="column" style="background-color:#ccc; border-style: solid; border-color: coral;">
    	<p>%s</p>
  	</div>
`

const Html = `
<!DOCTYPE html>
<html>
<head>
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
<style>
* {
  box-sizing: border-box;
}

/* Create two equal columns that floats next to each other */
.column {
  float: left;
  width: 100px;
  height: 100px;
  padding: 1px;
}

/* Clear floats after the columns */
.row:after {
  content: "";
  display: table;
  clear: both;
  border-color: coral;
}

</style>
</head>
<body>
	%s
</body>
</html>
`