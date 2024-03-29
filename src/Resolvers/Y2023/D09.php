<?php

declare(strict_types=1);

namespace App\Resolvers\Y2023;

use App\DTO\Solution;
use App\Resolvers\ResolverInterface;

class D09 implements ResolverInterface
{
    public function resolve(array $input): Solution
    {
        $totalOne = 0;
        $totalTwo = 0;

        foreach ($input as $row) {
            if (empty($row)) {
                continue;
            }

            $matches = [];
            preg_match_all('/-?\d+/', $row, $matches);
            $numbers = array_map('\intval', current($matches));
            $latests = [];
            $firsts = [];

            while ($this->notAllZero($numbers)) {
                $nbNumbers = \count($numbers);
                $firsts[] = $numbers[0];

                for ($x = 0; $x < $nbNumbers - 1; $x++) {
                    $numbers[$x] = $numbers[$x + 1] - $numbers[$x];
                }

                $latests[] = $numbers[$nbNumbers - 1];
                unset($numbers[$nbNumbers - 1]);
            }

            $totalOne += array_sum($latests);
            $totalTwo += $this->calculatePartTwo($firsts);
        }

        return new Solution($totalOne, $totalTwo);
    }

    private function notAllZero(array $numbers): bool
    {
        $count = array_count_values($numbers);

        return 1 !== \count($count) || 0 !== current(array_keys($count));
    }

    private function calculatePartTwo(array $numbers): int
    {
        $numbers[] = 0;
        $result = 0;
        $nbNumbers = \count($numbers);

        for ($x = $nbNumbers - 1; $x > 0; $x--) {
            $numbers[$x - 1] -= $numbers[$x];
            $result += $numbers[$x - 1] - $numbers[$x];
        }

        return $result;
    }
}
