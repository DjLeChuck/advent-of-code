<?php

declare(strict_types=1);

namespace App\Resolvers\Y2022;

use App\DTO\Solution;
use App\Resolvers\ResolverInterface;

class D03 implements ResolverInterface
{
    public function resolve(array $input): Solution
    {
        // First part
        $firstSumPriorities = 0;

        foreach ($input as $row) {
            if (empty($row)) {
                continue;
            }

            $chars = str_split($row);
            $count = \count($chars);
            $firstCompartment = array_slice($chars, 0, $count / 2);
            $secondCompartment = array_slice($chars, $count / 2);

            $this->sumPriorities($firstSumPriorities, $firstCompartment, $secondCompartment);
        }

        // Second part
        $secondSumPriorities = 0;

        while (!empty($input)) {
            // Group backs by three
            $group = array_splice($input, 0, 3);
            if (3 !== \count($group)) {
                continue;
            }

            // Split by chars on each group
            array_walk($group, static function (&$row) {
                $row = str_split($row);
            });

            $this->sumPriorities($secondSumPriorities, ...$group);
        }

        return new Solution($firstSumPriorities, $secondSumPriorities);
    }

    private function sumPriorities(int &$priorities, ...$array): void
    {
        $intersection = array_unique(array_intersect(...$array));

        foreach ($intersection as $char) {
            // Convert char to lower, then integer between 1 and 26
            $priorities += ord(mb_strtolower($char)) - 96;

            // If uppercase letter, add 26
            if (preg_match('`^[A-Z]$`', $char)) {
                $priorities += 26;
            }
        }
    }
}
