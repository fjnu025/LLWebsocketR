// This Header is auto generated by BDSLiteLoader Toolchain
#pragma once
#define AUTO_GENERATED
#include "../Global.h"
#include "Packet.hpp"

#define BEFORE_EXTRA
// Include Headers or Declare Types Here

#undef BEFORE_EXTRA

class EducationSettingsPacket : public Packet {

#define AFTER_EXTRA
// Add Member There

#undef AFTER_EXTRA

#ifndef DISABLE_CONSTRUCTOR_PREVENTION_EDUCATIONSETTINGSPACKET
public:
    class EducationSettingsPacket& operator=(class EducationSettingsPacket const&) = delete;
    EducationSettingsPacket(class EducationSettingsPacket const&) = delete;
#endif

public:
    /*0*/ virtual ~EducationSettingsPacket();
    /*1*/ virtual enum MinecraftPacketIds getId() const;
    /*2*/ virtual std::string getName() const;
    /*3*/ virtual void write(class BinaryStream&) const;
    /*4*/ virtual bool disallowBatching() const;
    /*5*/ virtual enum StreamReadResult _read(class ReadOnlyBinaryStream&);
    /*
    inline  ~EducationSettingsPacket(){
         (EducationSettingsPacket::*rv)();
        *((void**)&rv) = dlsym("??1EducationSettingsPacket@@UEAA@XZ");
        return (this->*rv)();
    }
    */
    MCAPI EducationSettingsPacket(struct EducationLevelSettings);
    MCAPI EducationSettingsPacket();

protected:

private:

};