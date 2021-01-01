#include "zig.h"

extern EB_Error_Code hookCallback(
	EB_Book* book,
	EB_Appendix* appendix,
	void* container,
	EB_Hook_Code hookCode,
	int argc,
	const unsigned int argv[]
);

EB_Error_Code installHook(EB_Hookset* hookset, EB_Hook_Code hookCode) {
    const EB_Hook hook = {hookCode, hookCallback};
    return eb_set_hook(hookset, &hook);
}
