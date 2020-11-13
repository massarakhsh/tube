var DT_Instant;

function grid_redraw(elm) {
    let path = elm.attr('path');
    front_proc(path, grid_redraw_init, elm);
}

function grid_redraw_init(elm, lika) {
    let instant = ('grid' in lika) ? lika.grid : [];
    grid_prepare(instant);
    let grid = elm.DataTable(instant);
    DT_Instant = instant;
    grid.on( 'select', grid_select);
}

function grid_prepare(data) {
    if (data !== null && typeof(data) == 'object') {
        for (var key in data) {
            let value = data[key];
            if (typeof(value) == 'string') {
                var match;
                if (match = /^function_(.+)\((.*)\)/.exec(value)) {
                    let func = match[1];
                    let parm = match[2];
                    if (func in window) {
                        data[key] = function () {
                            window[func](this, parm);
                        };
                    } else {
                        data[key] = grid_nothing;
                    }
                } else if (match = /^function_(.+)/.exec(value)) {
                    let func = match[1];
                    if (func in window) {
                        data[key] = window[func];
                    } else {
                        data[key] = grid_nothing;
                    }
                }
            } else if (value !== null && typeof(value) == 'object') {
                grid_prepare(data[key]);
            }
        }
    }
}

function grid_nothing() {
}

function grid_select( e, dt, type, indexes ) {
    if ( type === 'row' ) {
        if (indexes && indexes.length > 0) {
            var datas = dt.rows(indexes).data();
            if (datas && datas.length>0) {
                let name = datas[0].file;
                let path = DT_Instant.ajax;
                if (match = /^(.+)griddata(.*)$/.exec(path)) {
                    path = match[1] + "select/" + string_to_XS(name) + match[2];
                    get_data_part(path);
                }
            }
        }
    }
}

