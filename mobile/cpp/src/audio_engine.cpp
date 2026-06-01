#include "audio_engine.h"
#include <iostream>

void init_audio_engine() {
    std::cout << "Audio Engine Initialized" << std::endl;
}

void start_capture() {
    std::cout << "Audio Capture Started" << std::endl;
}

void stop_capture() {
    std::cout << "Audio Capture Stopped" << std::endl;
}

void destroy_audio_engine() {
    std::cout << "Audio Engine Destroyed" << std::endl;
}
