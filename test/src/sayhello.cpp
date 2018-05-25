#include "sayhello/sayhello.hpp"

#include <sstream>

SayHello::SayHello(const std::string& name)
    : m_name(name)
{
}

std::string SayHello::hello()
{
    std::stringstream stream;
    stream << "Hello " << m_name << "!";
    return stream.str();
}

std::string SayHello::goodbye()
{
    std::stringstream stream;
    stream << "Goodbye " << m_name << "!";
    return stream.str();
}

const std::string& SayHello::name()
{
    return m_name;
}

std::string* SayHello::mutable_name()
{
    return &m_name;
}


void SayHello::set_name(const std::string& name)
{
   m_name = name;
}
