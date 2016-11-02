<!-- Latest compiled and minified CSS -->

<html>
<head>
  <title>Mandlebrot Set Gif Maker</title>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link rel="stylesheet" href="http://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css">
  <script src="https://ajax.googleapis.com/ajax/libs/jquery/1.12.4/jquery.min.js"></script>
  <script src="http://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
</head>
<body>
<div class="jumbotron" style="text-align: center">
  <h1>Mandlebrot Set Gif Maker!</h1>
  <p>This is a simple gif generator for creating a gif of the convergence of the mandlebrot set</p>  
</div>
<div class="container" style="max-width: 600px">
<form action="/" method="post">
    <div class="form-group">
        <label for="minX">Minimum X value :</label>
        <input type="text" class="form-control" name="minX" placeholder="-2">
    </div>
    
    <div class="form-group">
        <label for="maxX">Maximum X value :</label>
        <input type="text" class="form-control" name="maxX" placeholder="1">
    </div>

    <div class="form-group">
        <label for="minY">Minimum Y value :</label>
        <input type="text" class="form-control" name="minY" placeholder="-1">
    </div>

    <div class="form-group">
        <label for="maxY">Maximum Y value :</label>
        <input type="text" class="form-control" name="maxY" placeholder="1">
    </div>
    
    <div class="form-group">
        <label for="iter">Number of Iterations :</label>
        <input type="text" class="form-control" name="iter" placeholder="10">
    </div>
    
    <div class="form-group">
        <label for="resl">Resolution :</label>
        <input type="text" class="form-control" name="resl" placeholder="100">
    </div>    
    
    <div><input type="submit" value="Generate"></div>
</form>
</div>
</body>
</html>