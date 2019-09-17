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
	<div class="column" style="background-color:#58ACFA; border-style: solid; border-color: coral;">
    	<p>%s</p>
  	</div>
`

const PlayerStatesList = `
<div class="block">
	<div class="row">
		<p>Player: Score </p>
	</div>
	%s
</div>
`

const PlayerState = `
	<div class="row">
		<p>%s: %d</p>
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

.left {
  float: left;
  width: 50%%;
}

.right {
  float: right;
  width: 50%%;
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

<div class="left">
%s
</div>

<div class="right">
%s
</div>
</body>
</html>
`
