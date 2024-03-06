const url = 'https://jsonplaceholder.typicode.com/todos';

interface Todo {
	userId : number;
	id : number;
	title : string;
	completed: boolean;
}

let printTable = (src : Todo[]) => {
  let limit : number = 10;
  let ident : string = '  ';
  console.log(`${ident}Id\tTitle\t\tCompleted`);
  console.log(`${ident}--------------------------------`);
  for (let todo of src) {
	limit--;
	if (limit == 0) {
	  break;
	}
	if (todo.title[9] == ' ') {
	console.log(`${ident}${todo.id}\t${todo.title.substring(0, 9)}... \t${todo.completed}`);
	} else {
	  console.log(`${ident}${todo.id}\t${todo.title.substring(0, 10)}...\t${todo.completed}`);
	}
  }
  if (limit == 0) {
	console.log(`${ident}\t<< ---- >>`);
	let todo : Todo = src[src.length -1] as Todo;

	if (todo.title[9] == ' ') {
	  console.log(`${ident}${todo.id}\t${todo.title.substring(0, 9)}... \t${todo.completed}`);
	} else {
	  console.log(`${ident}${todo.id}\t${todo.title.substring(0, 10)}...\t${todo.completed}`);
	}
  }
}

fetch(url).then(async (r) => {
	let k : string = await r.text() as string;
	let obj = JSON.parse(k) as Todo[];

	console.log(obj.length,'records returned');
	printTable(obj);
});
