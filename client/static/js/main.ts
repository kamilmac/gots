/// <reference path="../../typings/browser.d.ts"/>
import World from './modules/meow';
import {Dog as gog} from './modules/ruff';

class Hello {
    say() {
        console.log('hello');
    }
}

let h = new Hello;
let w = new World;
let d = new gog;

h.say();
w.shout();
d.bark();
console.log(Vue);
