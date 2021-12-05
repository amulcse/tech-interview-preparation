<?php
/*
 * Complete the 'sortedSum' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY a as parameter.
 */

$a = [1,2,3];

echo sortedSum($a);

function sortedSum($a) {
    // Write your code here
    $sortedA = $a;
    sort($sortedA);
    $sortedSum = 0;

    $currentSum = 0;
    $k=1;
    $sortedIndexLocation = [];

    $sum = 0;
    for($j=sizeof($sortedA)-1;$j>=0;$j--){
      $sum += $sortedA[$j];
      $sortedIndexLocation[$j] = $sum;
    }

    // $sortedSum = $currentSum;
    $removedSum = 0;

    $newArray = $sortedA;
    for($i = sizeof($a)-1; $i >= 0 ; $i--){
        $k=1;
        $ignoreElementDuringSum = [];
        $unsetDone = false;
        $sortedA = $newArray;
        $newArray = [];
        foreach($sortedA as $j => $value){
            $currentSum += $k*$sortedA[$j];
            $k++;
            if(!$unsetDone && $sortedA[$j] == $a[$i]){
                unset($sortedA[$j]); 
                $unsetDone = true;
            }
            else{
                $newArray[] = $sortedA[$j];
            }
            // if(!isset($igonoreElement[$sortedA[$j]]) || (isset($ignoreElementDuringSum[$sortedA[$j]]) && $ignoreElementDuringSum[$sortedA[$j]] >= $igonoreElement[$sortedA[$j]])){
            //     $currentSum += $k*$sortedA[$j];
            //     $k++;
            // }
            // else{
                
            //     if(!isset($ignoreElementDuringSum[$sortedA[$j]])){
            //         $ignoreElementDuringSum[$sortedA[$j]]= 0;
            //     }
            //     $ignoreElementDuringSum[$sortedA[$j]]++;
            // }
        }
        // $sortedSum[] = $currentSum;
        
        if(!isset($igonoreElement[$a[$i]])){
            $igonoreElement[$a[$i]]= 0;
        }
        $igonoreElement[$a[$i]]++;
    }
    
    return $currentSum%(1000000007);

}