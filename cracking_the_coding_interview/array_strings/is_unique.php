<?php

$string = 'amul';
$stringHasCharacters = [];
$unique = true;

if(strlen($string) > 26){
  $unique = false;
}
else{
  for($i = 0; $i < strlen($string) ; $i++){
      if(!isset($stringHasCharacters[$string[$i]])){
        $stringHasCharacters[$string[$i]] = 0;
      }
      else {
        $unique = false;
        break;
      }
  }
}

if($unique){
  echo $string." has unique characters.";
}
else{
  echo $string." has not unique characters.";
}