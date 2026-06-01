cmake_minimum_required(VERSION 3.10.2)

project(audio_engine VERSION 1.0.0 LANGUAGES CXX)

# Add Oboe
include(FetchContent)
FetchContent_Declare(
  oboe
  GIT_REPOSITORY https://github.com/google/oboe
  GIT_TAG 1.9.3
)
FetchContent_MakeAvailable(oboe)

# Use absolute paths with quotes to handle spaces in project path
add_library(audio_engine SHARED
    "${CMAKE_CURRENT_SOURCE_DIR}/../../../../../cpp/src/audio_engine.cpp"
)

target_include_directories(audio_engine PUBLIC
    "${CMAKE_CURRENT_SOURCE_DIR}/../../../../../cpp/include"
)

# Link Oboe and Android log library
target_link_libraries(audio_engine
    PUBLIC
    oboe
    log
)
