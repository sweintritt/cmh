cmake_minimum_required(VERSION 3.5)

project(sayhello VERSION 0.1.0 LANGUAGES CXX)

include_directories(${CMAKE_CURRENT_LIST_DIR}/include)

set(SOURCES ${CMAKE_CURRENT_LIST_DIR}/src/sayhello.cpp)

add_library(${PROJECT_NAME} ${SOURCES})

install(TARGETS ${PROJECT_NAME} DESTINATION lib)

install(FILES ${CMAKE_CURRENT_SOURCE_DIR}/include/sayhello/sayhello.hpp DESTINATION include/sayhello)
