#include <mutex>
#include <iostream>
using namespace std;

class Iphone
{
private:
    Iphone() = default;
    ~Iphone() = default;
    Iphone(const Iphone &nitin) = delete;
    Iphone operator=(const Iphone &nitin) = delete;
    
static Iphone *obj;
static mutex mutexobj;
public:
    static Iphone *getInstance(){
        
        if (nullptr == obj){
            lock_guard<mutex> lock(mutexobj);
            if(nullptr == obj){
                cout<<endl<<"getInstance---->";
                obj = new Iphone();
                lock_guard<mutex> unlock(mutexobj);
            }
        }
        return obj;
    }
    
    void f1(){
        cout<<endl<<"F1 is calling";
    }
    
    void f2(){
        cout<<endl<<"F2 is calling";
    }

};

Iphone *Iphone::obj=nullptr;
mutex Iphone::mutexobj;

int main(){
    for(int i=0; i<10;i++){
    
    // Iphone *obj = new Iphone();
    // obj->f1();
    Iphone::getInstance()->f1();
    Iphone::getInstance()->f2();
    }
    return 0;
}