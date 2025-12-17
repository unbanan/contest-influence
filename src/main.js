import Two from 'two.js';

import {Field, Cell, Player, TickView, Visualization} from './utils';

const container = document.getElementById('game-canvas');

const params = {
    autostart: true,
    height: container.clientHeight,
    width: container.clientWidth,
    type: Two.Types.svg,
};
const two = new Two(params).appendTo(container);

const R = 30;
const r = Math.sqrt(R ** 2 - (R / 2) ** 2)
const sides = 6;
const padding = 5;
const min_l = 20, max_l = 50;
const max_value = 8;

// const w_n = 10;
// const h_n = Math.floor((container.clientHeight - r) / r);
const field = two.makeGroup()

for (let i = 0; ; i++) {
    const y = padding + r + i * r + i * padding / 2;
    if (y + r + padding >= container.clientHeight) {
        break;
    }
    for (let j = 0; ; j++) {
        if (Math.random() > 0.3) {

            const x = padding + R + j * 3 * R + 2 * padding * j + (i % 2) * padding + R * 1.5 * (i % 2);
            if (x + R + padding >= container.clientWidth) {
                break;
            }

            const hexagon = two.makePolygon(0, 0, R, sides);
            hexagon.fill = 'rgba(171, 171, 171, 0.04)';
            hexagon.linewidth = 0;

            const inner_hexagon = two.makePolygon(0, 0, r * 0.9, sides);
            inner_hexagon.fill = 'rgba(54, 48, 48, 1)';
            hexagon.id = `hexagon_id=${i}_${j}`;
            inner_hexagon.rotation = Math.PI / 6;
            inner_hexagon.id = `inner_hexagon_id=${i}_${j}`
            inner_hexagon.linewidth = 1;
            two.update();
            inner_hexagon._renderer.elem.classList.add("untouchable");
            
            var value = 0;
            var value = Math.random() > 0.5 ? Math.ceil(max_value * Math.random()) : 0;
            const ratio = Math.min(value / max_value, 1);
            const light = max_l - (ratio * (max_l - min_l));
            
            if (value !== 0) {
                inner_hexagon.fill = `hsl(200, 30%, ${light}%)`;
            }
            
            
            const text = two.makeText(`${value}`, 0, 0);
            text.fill = value !== 0 ? 'white' : 'transparent';
            text.size = 12;
            text.weight = 700;
            text.family = 'sans-serif';
            
            const group = two.makeGroup();
            const hg = two.makeGroup(hexagon, inner_hexagon);
            group.add(hg, text);
            group.translation.set(x, y);
            field.add(group)
            two.update();
        } else {
        }
    }
}

setupInteractions(two);

function setupInteractions(instance) {

    var result = new Set()

    const svg = instance.renderer.domElement;

    const getShape = (e) => {
        if (e.target.tagName === 'path') {
            return instance.scene.getById(e.target.id);
        }
        return null;
    };

    svg.addEventListener('click', (e) => {
        const shape = getShape(e);
        if (shape) {
            if (shape.id) {
                if (result.has(shape.id)) {
                    result.delete(shape.id);
                } else {
                    result.add(shape.id);
                }
            } 
            if (shape.stroke == "") {
                shape.noStroke();
            } else {
                shape.stroke = '#f5f6fa';
            }
            if (!shape.linewidth) {
                shape.linewidth = 4;
            } else {
                shape.linewidth = 0;
            }
        }
    });
}

two.update();