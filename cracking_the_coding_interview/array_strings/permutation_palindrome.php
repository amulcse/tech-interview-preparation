<?php

$string = "ababc";

echo $string." has a permuation of palindrome - ".hasPermutationPalindrome($string);

function hasPermutationPalindrome($string){

  $characterCounter = [] ;
  $total = 0;
  for($i = 0;$i < strlen($string) ; $i++){
    if($string[$i] !== ''){
      if(!isset( $characterCounter[$string[$i]])){
        $characterCounter[$string[$i]] = 0;
      }
      $characterCounter[$string[$i]] = ($characterCounter[$string[$i]]+1)%2; 
      if($characterCounter[$string[$i]] == 1){
        $total++;
      }
      else{
        $total--;
      }
    }
  }

  if($total > 1){
    return 'No';
  }
  else{
    return 'Yes';
  }

}