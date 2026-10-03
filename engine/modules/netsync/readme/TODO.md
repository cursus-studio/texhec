This module needs clean up for transparent events and ticks.
Currently we hope event won't land on tick instead of handling it properly.
Transparent events aren't tracked to reverse back in state.
Client also hopes he doesn't receive event at a bad time because if event from server lands in a bad time
transparent events are going to be messed.


