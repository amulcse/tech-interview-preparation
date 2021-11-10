<?php

/*
 * Complete the 'climbingLeaderboard' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY ranked
 *  2. INTEGER_ARRAY player
 */

function climbingLeaderboard($ranked, $player) {
    // Write your code here
    $result = [];
    
    $i = sizeof($player) - 1;
    $currentRank = 1;
    for($j = 0; $j < sizeof($ranked) ; $j++){
        if($j > 0 && $ranked[$j] < $ranked[$j-1] ){
            $currentRank++;    
        }
        while($i >=0 && $player[$i] >= $ranked[$j]){
            $result[$i] = $currentRank;
            $i--;   
        }
        if($i < 0){
            break;
        }
    }
    $currentRank++;
    $remainingRank = sizeof($player) - sizeof($result);
    for($i = 0; $i < $remainingRank  ; $i++){
        $result[$i] = $currentRank;
    }
    $finalResult = [];
    for($i = 0; $i < sizeof($result) ;$i++){
        $finalResult[] = $result[$i];
    }
    return $finalResult;

}

$fptr = fopen(getenv("OUTPUT_PATH"), "w");

$ranked_count = intval(trim(fgets(STDIN)));

$ranked_temp = rtrim(fgets(STDIN));

$ranked = array_map('intval', preg_split('/ /', $ranked_temp, -1, PREG_SPLIT_NO_EMPTY));

$player_count = intval(trim(fgets(STDIN)));

$player_temp = rtrim(fgets(STDIN));

$player = array_map('intval', preg_split('/ /', $player_temp, -1, PREG_SPLIT_NO_EMPTY));

$result = climbingLeaderboard($ranked, $player);

fwrite($fptr, implode("\n", $result) . "\n");

fclose($fptr);
<?php

/*
 * Complete the 'climbingLeaderboard' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY ranked
 *  2. INTEGER_ARRAY player
 */

function climbingLeaderboard($ranked, $player) {
    // Write your code here
    $result = [];
    
    $i = sizeof($player) - 1;
    $currentRank = 1;
    for($j = 0; $j < sizeof($ranked) ; $j++){
        if($j > 0 && $ranked[$j] < $ranked[$j-1] ){
            $currentRank++;    
        }
        while($i >=0 && $player[$i] >= $ranked[$j]){
            $result[$i] = $currentRank;
            $i--;   
        }
        if($i < 0){
            break;
        }
    }
    $currentRank++;
    $remainingRank = sizeof($player) - sizeof($result);
    for($i = 0; $i < $remainingRank  ; $i++){
        $result[$i] = $currentRank;
    }
    $finalResult = [];
    for($i = 0; $i < sizeof($result) ;$i++){
        $finalResult[] = $result[$i];
    }
    return $finalResult;

}

$fptr = fopen(getenv("OUTPUT_PATH"), "w");

$ranked_count = intval(trim(fgets(STDIN)));

$ranked_temp = rtrim(fgets(STDIN));

$ranked = array_map('intval', preg_split('/ /', $ranked_temp, -1, PREG_SPLIT_NO_EMPTY));

$player_count = intval(trim(fgets(STDIN)));

$player_temp = rtrim(fgets(STDIN));

$player = array_map('intval', preg_split('/ /', $player_temp, -1, PREG_SPLIT_NO_EMPTY));

$result = climbingLeaderboard($ranked, $player);

fwrite($fptr, implode("\n", $result) . "\n");

fclose($fptr);
