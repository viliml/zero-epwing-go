#include "zig.h"

extern EB_Error_Code hookCallback(
	EB_Book* ebBook,
	EB_Appendix* ebAppendix,
	void* container,
	EB_Hook_Code ebHookCode,
	int argc,
	const unsigned int argv[]
);

EB_Error_Code installHook(EB_Hookset* ebHookset, EB_Hook_Code ebHookCode) {
    const EB_Hook ebHook = {ebHookCode, hookCallback};
    return eb_set_hook(ebHookset, &ebHook);
}
