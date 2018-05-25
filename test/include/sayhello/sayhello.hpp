
#ifndef SAY_HELLO_INCLUDE_SAY_HELLO_HPP_
#define SAY_HELLO_INCLUDE_SAY_HELLO_HPP_

#include <string>

class SayHello
{
    public:
        SayHello(const std::string& name);

        std::string hello();

        std::string goodbye();

        const std::string& name();

        std::string* mutable_name();

        void set_name(const std::string& name);

    protected:
        std::string m_name;
};


#endif /* SAY_HELLO_INCLUDE_SAY_HELLO_HPP_ */
