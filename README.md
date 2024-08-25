An extremely short table-driven parser that can be used for parallel & incremental parsing.

This uses generalized operator precedence parsing as described in https://www.sciencedirect.com/science/article/pii/S0167642315002610
, where the parser input can be arbitrary nonterminals, not just terminals. It has the property that you can
run several shift reduce parsers on different segments of the tokenized input, and merge them by concatenating their
shift reduce stacks and passing them as input to the same parser, which enables parallel or incremental parsing.
