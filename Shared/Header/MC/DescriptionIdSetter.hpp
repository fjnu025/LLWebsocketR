// This Header is auto generated by BDSLiteLoader Toolchain
#pragma once
#define AUTO_GENERATED
#include "../Global.h"

#define BEFORE_EXTRA
// Include Headers or Declare Types Here

#undef BEFORE_EXTRA

class DescriptionIdSetter {

#define AFTER_EXTRA
// Add Member There

#undef AFTER_EXTRA

#ifndef DISABLE_CONSTRUCTOR_PREVENTION_DESCRIPTIONIDSETTER
public:
    class DescriptionIdSetter& operator=(class DescriptionIdSetter const&) = delete;
    DescriptionIdSetter(class DescriptionIdSetter const&) = delete;
    DescriptionIdSetter() = delete;
#endif

public:
    MCAPI bool Deserialize(class BasicLoader&, struct SerializerTraits const&, class BedrockLoadContext const&);
    MCAPI bool Serialize(class BasicSaver&, struct SerializerTraits const&) const;

protected:

private:

};