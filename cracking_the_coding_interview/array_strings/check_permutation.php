<?php

$string1 = 'abc';
$string2 = 'bac';

$isPermutaion = true;

$string1CharacterCounts = [];
$string2CharacterCounts = [];

/* if both string length is not matched */
if(strlen($string1) !== strlen($string2)){
  $isPermutaion = false;
}
else{
  /* check each character occurance count */
  for($i = 0; $i < strlen($string2) ; $i++){

    if(!isset($string1CharacterCounts[$string1[$i]])){
      $string1CharacterCounts[$string1[$i]] = 1;
    }
    else{
      $string1CharacterCounts[$string1[$i]]++;
    }

    if(!isset($string2CharacterCounts[$string2[$i]])){
      $string2CharacterCounts[$string2[$i]] = 1;
    }
    else{
      $string2CharacterCounts[$string2[$i]]++;
    }

  }

  if(sizeof($string1CharacterCounts) !== sizeof($string2CharacterCounts)){
    $isPermutaion = false;
  }
  else{
    foreach($string1CharacterCounts as $character => $count){
      if($count !== $string2CharacterCounts[$character]){
        $isPermutaion = false;
        break;
      }
    }
  }
}

if($isPermutaion){
  echo $string2." is the permutaion of ".$string1."!";
}
else{
  echo $string2." is not the permutaion of ".$string1."!";
}