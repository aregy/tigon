const insort = function (arr){
	let d = 0;
	for (let i = 0; i < arr.length; i++){
		if (arr[i] == 0){
			d++;
			continue;
		}
		if (d != 0){
			arr[i-d] = arr[i];
		}
		if (i >= arr.length - 1 - d){
			arr[i] = 0;
		}
	}
};

module.exports = { insort }; 
